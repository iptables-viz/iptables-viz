import KubernetesHeader from "../components/KubernetesHeader";
import LinuxHeader from "../components/LinuxHeader";
import React from "react";

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
