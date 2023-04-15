import CytoscapeComponent from "react-cytoscapejs";
import { makeTippy } from "../utils/Tippy";
import { cytoscapeStylesheet } from "../utils/CytoscapeStylesheet";
import Cytoscape from "cytoscape";
import popper from "cytoscape-popper";
import dagre from "cytoscape-dagre";

Cytoscape.use(popper);
Cytoscape.use(dagre);

export default function Graph(props: {
  nodes: Cytoscape.ElementDefinition[];
  edges: Cytoscape.ElementDefinition[];
}) {
  return (
    <CytoscapeComponent
      elements={CytoscapeComponent.normalizeElements({
        nodes: props.nodes,
        edges: props.edges,
      })}
      stylesheet={cytoscapeStylesheet}
      style={{
        width: `98vw`,
        height: `80vh`,
      }}
      cy={(cy: Cytoscape.Core) => {
        cy.on("mouseover", "edge", (e: Cytoscape.EventObjectEdge) => {
          let edge = e.target;
          edge.on("mouseover", (_) => {
            let tippy = makeTippy(edge, edge.data("label"));
            tippy.show();
            edge.on("mouseout", (_) => {
              tippy.destroy();
            });
          });
        });

        cy.on("mouseover", "node", (e: Cytoscape.EventObjectNode) => {
          let node = e.target;

          let tippy = makeTippy(node, node.data("name"));
          tippy.show();
          setTimeout(() => tippy.destroy(), 2000);

          cy.elements()
            .difference(node.outgoers())
            .not(node)
            .addClass("semitransp");
          node.addClass("highlight").outgoers().addClass("highlight");

          node.on("mouseout", () => {
            tippy.destroy();

            cy.elements().removeClass("semitransp");
            node.removeClass("highlight").outgoers().removeClass("highlight");
          });
        });

        cy.on("add remove", "node", () => {
          cy.layout({
            name: "dagre",
          }).run();
        });
      }}
    />
  );
}
