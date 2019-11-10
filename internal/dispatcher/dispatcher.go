package dispatcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
)

var defaultSize = 100
var searchURL = "http://registry.npmjs.org/-/v1/search?text=not:unstable&quality=0.0&maintenance=0.0&popularity=1.0&size=%d&from=%d"

type objects struct {
	Objects []*multiPackageInfo `json:"objects,omitempty"`
}

type multiPackageInfo struct {
	Package *packageInfo `json:"package"`
}

type packageInfo struct {
	Name    string `json:name`
	Version string `json:version`
}

type Dispatcher struct {
	q         queue.Queue
	bootstrap bool
	topN      int
}

func NewDispatcher(q queue.Queue, bootstrap bool, topN int) *Dispatcher {
	return &Dispatcher{q: q, bootstrap: bootstrap, topN: topN}
}

func (d *Dispatcher) Start(ctx context.Context) error {
	if !d.bootstrap {
		return nil
	}

	counter := 0
	for counter < d.topN {
		pageSize := d.calcPageSize(counter, defaultSize, d.topN)
		endpoint := fmt.Sprintf(searchURL, pageSize, counter)
		fmt.Printf("calling endpoint %s", endpoint)
		resp, err := http.Get(endpoint)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("not found")
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("non-200 status code")
		}

		objs := &objects{}
		if err := json.NewDecoder(resp.Body).Decode(objs); err != nil {
			return err
		}

		for _, p := range objs.Objects {
			d.q.Add(ctx, fmt.Sprintf("%s@%s", p.Package.Name, p.Package.Version))
		}

		counter += pageSize
	}

	return nil
}

func (d *Dispatcher) calcPageSize(counter, defaultPageSize, maxN int) int {
	if counter+defaultPageSize > maxN {
		return maxN - counter
	}

	return defaultPageSize
}
