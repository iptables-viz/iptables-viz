import KubernetesHeader from "../components/KubernetesHeader";
import LinuxHeader from "../components/LinuxHeader";
import React from "react";

/**
 * Header container for specifying the parameters for the iptables output.
 *
 * @param TableType - Type of the table selected
 * @param SetTableType - UseState hook setter for TableType
 * @param PodName - Name of the pod selected
 * @param SetPodName - UseState hook setter for PodName
 * @param PodNames - Name of the kube-proxy pods in the kube-system namespace
 *
 * @returns The JSX header component definition.
 *
 * @beta
 */
export const Header = (props: {
  TableType: string;
  SetTableType: React.Dispatch<React.SetStateAction<string>>;
  PodName: string;
  SetPodName: React.Dispatch<React.SetStateAction<string>>;
  PodNames: string[];
}): JSX.Element => {
  switch (import.meta.env.VITE_PLATFORM) {
    case "kubernetes":
      return (
        <KubernetesHeader
          TableType={props.TableType}
          SetTableType={props.SetTableType}
          PodName={props.PodName}
          SetPodName={props.SetPodName}
          PodNames={props.PodNames}
        />
      );
    default:
      return (
        <LinuxHeader
          TableType={props.TableType}
          SetTableType={props.SetTableType}
        />
      );
  }
};
