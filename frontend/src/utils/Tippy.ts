import Cytoscape from "cytoscape";
import tippy from "tippy.js";

export const makeTippy = (
  ele: Cytoscape.EdgeSingular | Cytoscape.NodeSingular,
  label: String
) => {
  let ref = ele.popperRef();
  let dummyDomEle = document.createElement("div");

  return tippy(dummyDomEle, {
    getReferenceClientRect: ref.getBoundingClientRect,
    trigger: "manual",
    content: () => {
      let div = document.createElement("div");
      label.split(",").forEach((field) => {
        let d = document.createElement("div");
        d.style.overflowWrap = "break-word";
        d.style.fontSize = "16px";
        d.appendChild(document.createTextNode(field));
        div.appendChild(d);
      });
      return div;
    },
    arrow: true,
    placement: "bottom",
    hideOnClick: false,
    sticky: "reference",
    interactive: true,
    appendTo: document.body,
  });
};
