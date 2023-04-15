import { GetKubernetesIptables } from "./KubernetesAPI";
import { GetLinuxIptables } from "./LinuxAPI";

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
