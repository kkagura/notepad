import { ServiceResponse } from "@/hooks/useRequest";
import { GetDocTree } from "wailsjs/go/services/DocService";

type ReturnTypeHelper<T extends (...args: unknown[]) => unknown> =
  ReturnType<T> extends Promise<infer R> ? R : unknown;

export type DocTreeResponseDTO = ReturnTypeHelper<typeof GetDocTree>;
export const getDocTree = (): Promise<ServiceResponse<DocTreeResponseDTO>> => {
  return GetDocTree().then((data) => {
    const res: ServiceResponse<DocTreeResponseDTO> = {
      code: 200,
      message: "",
      success: true,
      data,
    };
    return res;
  });
};
