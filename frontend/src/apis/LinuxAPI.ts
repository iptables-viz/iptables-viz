import { GetErrorMessage } from "../utils/Common";

/**
 * Fetches the Linux system iptables data to be used for generating the DAG.
 *
 * @param setIptablesData - UseState hook setter for iptables data
 * @param tableType - iptables table type to be fetched
 * @param setError - UseState hook setter for fetch error
 *
 * @returns A promise for the fetch API call to the backend.
 */
export const GetLinuxIptables = async (
  setIptablesData: React.Dispatch<React.SetStateAction<string>>,
  tableType: string,
  setError: React.Dispatch<React.SetStateAction<string>>
): Promise<void> => {
  try {
    const res = await fetch(
      `${import.meta.env.VITE_BASE_URL}/iptables/linux/${tableType}`
    );
    const json = await res.json();
    setIptablesData(json.iptablesOutput);
  } catch (err) {
    setError(GetErrorMessage(err));
  }
};
