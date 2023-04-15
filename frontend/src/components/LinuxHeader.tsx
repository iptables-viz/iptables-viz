import { Center, Select } from "@chakra-ui/react";

export default function LinuxHeader(props: {
  TableType: string;
  SetTableType: React.Dispatch<React.SetStateAction<string>>;
}): JSX.Element {
  return (
    <Center>
      <Select
        color="black"
        w="40%"
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
  );
}
