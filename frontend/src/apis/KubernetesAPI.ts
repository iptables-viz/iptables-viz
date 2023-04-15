import { GetErrorMessage } from "../utils/Common";

export const GetKubernetesIptables = async (
  setIptablesData: React.Dispatch<React.SetStateAction<string>>,
  tableType: string,
  setError: React.Dispatch<React.SetStateAction<string>>,
  setPodName: React.Dispatch<React.SetStateAction<string>>,
  setKubeProxyPodNames: React.Dispatch<React.SetStateAction<string[]>>,
  podName: string
) => {
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
