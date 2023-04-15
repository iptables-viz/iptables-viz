import { GetErrorMessage } from "../utils/Common";

export const GetLinuxIptables = async (
  setIptablesData: React.Dispatch<React.SetStateAction<string>>,
  tableType: string,
  setError: React.Dispatch<React.SetStateAction<string>>
) => {
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
