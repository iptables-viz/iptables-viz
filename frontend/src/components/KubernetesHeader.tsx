import { Flex, Center, Select } from "@chakra-ui/react";

/**
 * Header for selecting Kubernetes pod names and iptables table type.
 *
 * @param TableType - Type of the table selected
 * @param SetTableType - UseState hook setter for TableType
 * @param PodName - Name of the pod selected
 * @param SetPodName - UseState hook setter for PodName
 * @param PodNames - Name of the kube-proxy pods in the kube-system namespace
 *
 * @returns The JSX Kubernetes header component definition.
 */
export default function KubernetesHeader(props: {
  TableType: string;
  SetTableType: React.Dispatch<React.SetStateAction<string>>;
  PodName: string;
  SetPodName: React.Dispatch<React.SetStateAction<string>>;
  PodNames: string[];
}): JSX.Element {
  return (
    <Flex>
      <Center w="50%">
        <Select
          color="black"
          value={props.PodName}
          w="70%"
          placeholder="Select Kube Proxy Pod"
          onChange={(e) => {
            e.preventDefault();
            props.SetPodName(e.target.value);
          }}
        >
          {props.PodNames.map((podName: string) => {
            return (
              <option key={podName} color="black" value={podName}>
                {podName}
              </option>
            );
          })}
        </Select>
      </Center>
      <Center w="50%">
        <Select
          color="black"
          w="70%"
          placeholder="Select Iptable Type"
          value={props.TableType}
          onChange={(e) => {
            e.preventDefault();
            props.SetTableType(e.target.value);
          }}
        >
          <option value="nat">nat</option>
          <option value="filter">filter</option>
          <option value="raw">raw</option>
          <option value="mangle">mangle</option>
        </Select>
      </Center>
    </Flex>
  );
}
