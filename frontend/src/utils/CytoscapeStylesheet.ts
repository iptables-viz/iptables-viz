const cytoscapeStylesheet: cytoscape.Stylesheet[] = [
  {
    selector: "edge",
    css: {
      "curve-style": "bezier",
      "target-arrow-shape": "triangle",
    },
  },
  {
    selector: "node",
    style: {
      label: "data(label)",
      "font-size": "10px",
      backgroundColor: "black",
    },
  },
  {
    selector: "node.highlight",
    style: {
      "border-color": "#FFF",
      "border-width": "2px",
    },
  },
  {
    selector: "node.semitransp",
    style: {
      opacity: 0.5,
    },
  },
  {
    selector: "edge.highlight",
    style: {
      "mid-target-arrow-color": "#FFF",
    },
  },
  {
    selector: "edge.semitransp",
    style: {
      opacity: 0.2,
    },
  },
];

export { cytoscapeStylesheet };
