import Cytoscape from "cytoscape";
import { Header } from "./containers/Header";
import { Heading, Center, Box } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { ParseIptablesData } from "./utils/Data";
import Graph from "./components/Graph";
import GetIptables from "./apis/API";
import { Alert, AlertIcon } from "@chakra-ui/react";

import "tippy.js/dist/tippy.css";

/**
 * Application container for executing the iptables visualization logic.
 *
 * @returns The JSX application container definition.
 */
export default function App(): JSX.Element {
  const [tableType, setTableType] = useState<string>("nat");
  const [kubeProxyPodNames, setKubeProxyPodNames] = useState<string[]>([]);
  const [podName, setPodName] = useState<string>("");
  const [iptablesData, setIptablesData] = useState<string>("");
  const [error, setError] = useState<string>("");

  const [nodes, setNodes] = useState<Cytoscape.ElementDefinition[]>([]);
  const [edges, setEdges] = useState<Cytoscape.ElementDefinition[]>([]);

  useEffect(() => {
    GetIptables(
      setIptablesData,
      tableType,
      setError,
      setPodName,
      setKubeProxyPodNames,
      podName
    );
  }, [tableType, podName]);

  useEffect(() => {
    if (iptablesData !== "") {
      setError("");
      const [tempNodes, tempEdges] = ParseIptablesData(iptablesData);
      setNodes(tempNodes);
      setEdges(tempEdges);
    }
  }, [iptablesData, tableType, podName]);

  return (
    <div className="App">
      {error && (
        <Box px={30} pt={30}>
          <Alert status="error" variant="subtle">
            <AlertIcon />
            {error}
          </Alert>
        </Box>
      )}
      <Box w="80vw" h="11vh" m="auto">
        <Center my="25px">
          <Heading>Iptables Visualizer</Heading>
        </Center>
        <Header
          TableType={tableType}
          SetTableType={setTableType}
          PodName={podName}
          SetPodName={setPodName}
          PodNames={kubeProxyPodNames}
        />
      </Box>
      <Box>
        <Graph nodes={nodes} edges={edges} />
      </Box>
    </div>
  );
}
