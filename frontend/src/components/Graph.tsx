import CytoscapeComponent from "react-cytoscapejs";
import { makeTooltip } from "../utils/Tooltip";
import { cytoscapeStylesheet } from "../utils/CytoscapeStylesheet";
import Cytoscape from "cytoscape";
import popper from "cytoscape-popper";
import dagre from "cytoscape-dagre";

Cytoscape.use(popper);
Cytoscape.use(dagre);

/**
 * Cytoscape graph component for rendering the iptables rules as a DAG.
 *
 * @param nodes - The list of nodes in the DAG
 * @param edges - The list of edges in the DAG
 *
 * @returns The JSX Cytoscape graph component definition.
 */
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
            let tooltip = makeTooltip(edge, edge.data("label"));
            tooltip.show();
            edge.on("mouseout", (_) => {
              tooltip.destroy();
            });
          });
        });

        cy.on("mouseover", "node", (e: Cytoscape.EventObjectNode) => {
          let node = e.target;

          let tooltip = makeTooltip(node, node.data("name"));
          tooltip.show();
          setTimeout(() => tooltip.destroy(), 2000);

          cy.elements()
            .difference(node.outgoers())
            .not(node)
            .addClass("semitransp");
          node.addClass("highlight").outgoers().addClass("highlight");

          node.on("mouseout", () => {
            tooltip.destroy();

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
