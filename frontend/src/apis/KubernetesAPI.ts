import { GetErrorMessage } from "../utils/Common";

/**
 * Fetches the Kuberentes pod iptables data to be used for generating the DAG.
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
export const GetKubernetesIptables = async (
  setIptablesData: React.Dispatch<React.SetStateAction<string>>,
  tableType: string,
  setError: React.Dispatch<React.SetStateAction<string>>,
  setPodName: React.Dispatch<React.SetStateAction<string>>,
  setKubeProxyPodNames: React.Dispatch<React.SetStateAction<string[]>>,
  podName: string
): Promise<void> => {
  try {
    let res: Response;
    if (podName !== "") {
      res = await fetch(
        `${
          import.meta.env.VITE_BASE_URL
        }/iptables/kubernetes/${podName}/${tableType}`
      );
      const json = await res.json();
      setIptablesData(json.iptablesOutput);
    } else {
      res = await fetch(`${import.meta.env.VITE_BASE_URL}/iptables/kubernetes`);
      const json = await res.json();
      setIptablesData(json.iptablesOutput);
      setKubeProxyPodNames?.(json.podNames);
      setPodName?.(json.podName);
    }
  } catch (err) {
    setError(GetErrorMessage(err));
  }
};
