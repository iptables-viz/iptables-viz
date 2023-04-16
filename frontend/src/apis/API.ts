import { GetKubernetesIptables } from "./KubernetesAPI";
import { GetLinuxIptables } from "./LinuxAPI";

/**
 * Fetches the iptables data to be used for generating the DAG.
 *
 * @param setIptablesData - UseState hook setter for iptables data
 * @param tableType - iptables table type to be fetched
 * @param setError - UseState hook setter for fetch error
 * @param setPodName - UseState hook setter for the kube-proxy pod whose iptables need to be fetched
 * @param setKubeProxyPodNames - UseState hook setter for kube-proxy pods
 * @param podName - Pod whose iptables need to be fetched
 *
 * @returns A promise for the fetch API call to the backend.
 */
export default function GetIptables(
  setIptablesData: React.Dispatch<React.SetStateAction<string>>,
  tableType: string,
  setError: React.Dispatch<React.SetStateAction<string>>,
  setPodName?: React.Dispatch<React.SetStateAction<string>>,
  setKubeProxyPodNames?: React.Dispatch<React.SetStateAction<string[]>>,
  podName?: string
): Promise<void> {
  switch (import.meta.env.VITE_PLATFORM) {
    case "kubernetes":
      if (
        typeof setPodName !== "undefined" &&
        typeof setKubeProxyPodNames !== "undefined" &&
        typeof podName !== "undefined"
      ) {
        return GetKubernetesIptables(
          setIptablesData,
          tableType,
          setError,
          setPodName,
          setKubeProxyPodNames,
          podName
        );
      }
    default:
      return GetLinuxIptables(setIptablesData, tableType, setError);
  }
}
