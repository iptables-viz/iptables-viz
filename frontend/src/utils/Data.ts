import Cytoscape from "cytoscape";
import { IptablesSchema } from "../types/Types";

const TrimString = (label: string) => {
  return label.length > 10 ? label.substring(0, 10) + "..." : label;
};

export const ParseIptablesData = (
  iptablesData: string
): [
  nodes: Cytoscape.ElementDefinition[],
  edges: Cytoscape.ElementDefinition[]
] => {
  let parsedData: IptablesSchema = JSON.parse(iptablesData);
  let tempNodes: Cytoscape.ElementDefinition[] = [];
  let tempEdges: Cytoscape.ElementDefinition[] = [];
  parsedData.forEach((chain) => {
    tempNodes.push({
      data: {
        id: chain.chain,
        label: TrimString(chain.chain),
        name: chain.chain,
      },
    });
    chain.rules.forEach((rule) => {
      tempNodes.push({
        data: {
          id: rule.target,
          label: TrimString(rule.target),
          name: rule.target,
        },
      });
      tempEdges.push({
        data: {
          source: chain.chain,
          target: rule.target,
          label: `prot: ${rule.prot ?? "''"},\n
            opt: ${rule.opt ?? "''"},\n
            source: ${rule.source ?? "''"},\n
            destination: ${rule.destination ?? "''"},\n
            options: ${rule.options ?? "''"}`,
        },
      });
    });
  });
  return [tempNodes, tempEdges];
};
