import { Center, Select } from "@chakra-ui/react";

/**
 * Header for selecting iptables table type.
 *
 * @param TableType - Type of the table selected
 * @param SetTableType - UseState hook setter for TableType
 *
 * @returns The JSX Linux header component definition.
 */
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
