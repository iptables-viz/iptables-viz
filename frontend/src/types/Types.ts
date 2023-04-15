export type IptablesSchema = {
  chain: string;
  rules: {
    num: number;
    pkts: number;
    bytes: number; // converted based on suffix
    target: string;
    prot: string;
    opt: string; // -- = Null
    in: string;
    out: string;
    source: string;
    destination: string;
    options: string;
  }[];
}[];
