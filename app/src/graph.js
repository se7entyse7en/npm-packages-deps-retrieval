import React from 'react';
import Graph from 'react-graph-vis';

class GraphViz extends React.Component {

  render = () => {
    const deps = this.props.dependencies;
    const viz = this.buildViz(deps);

    return (
      <div className="row">
        <div className="col-12">
          {viz}
        </div>
      </div>
    );
  }

  buildViz = (deps) => {
    const graph = {
      nodes: [],
      edges: []
    };
    const packages = {};
    const options = {
      nodes: {
        shape: "dot",
        chosen: {
          label: function(values, id, selected, hovering) {
            values.size = values.size * 2;
          }
        }
      },
      physics: {
        enabled: false
      },
      layout: {
        hierarchical: {
          nodeSpacing: 100
        }
      },
      edges: {
        color: "#000000",
        chosen: {
          edge: function(values, id, selected, hovering) {
            values.color = "#ff0000";
            values.width = values.width * 3;
          }
        }
      }
    };

    this.buildTree(deps, graph, packages, 0);

    return (
      <Graph
        graph={graph}
        options={options}
      />
    );
  };

  buildTree = (deps, graph, packages, level) => {
    const name = deps.getName();
    const version = deps.getVersion();
    const children = deps.getDependenciesList();

    const packageNodeId = this.addPackage(name, version, graph, packages, level);

    for (let child of children) {
      const childName = child.getName();
      const childVersion = child.getVersion();
      const childPackageNodeId = this.addPackage(childName, childVersion, graph, packages, level + 1);

      graph.edges.push({from: packageNodeId, to: childPackageNodeId});

      this.buildTree(child, graph, packages, level + 1);
    }
  };

  addPackage = (name, version, graph, packages, level) => {
    const key = name + "@" + version;
    if (!(key in packages)) {
      packages[key] = graph.nodes.length;
      graph.nodes.push({
        id: packages[key],
        label: key,
        level: level
      });
    }

    graph.nodes[packages[key]]["level"] = Math.max(
      graph.nodes[packages[key]]["level"], level);

    return packages[key];
  };

}

export default GraphViz;
