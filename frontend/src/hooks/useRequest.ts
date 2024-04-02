import { reactive } from "vue";

class RequestError extends Error {
  constructor(public code: number, public message: string) {
    super(message);
  }
}

export interface ServiceResponse<Res = any> {
  code: number;
  message: string;
  data: Res;
  success: boolean;
}

interface RequestContext<Res = any, Req = any> {
  loading: boolean;
  pending: boolean;
  error: RequestError | null;
  request: (req?: Req) => Promise<Res>;
}

export const useRequest = <Res = any, Req = any>(
  requestFn: (req?: Req) => Promise<ServiceResponse<Res>>
) => {
  let seed = 0;
  const tasks = new Set<number>();
  const context: RequestContext<Res, Req> = reactive({
    loading: false,
    pending: false,
    error: null,
    request(req?: Req) {
      context.loading = true;
      context.pending = true;
      seed++;
      tasks.add(seed);
      return requestFn(req)
        .then((res) => {
          if (res.success) {
            context.error = null;
            return res.data;
          }
          context.error = new RequestError(res.code, res.message);
          return Promise.reject(context.error);
        })
        .catch((error) => {
          if (context.error) return Promise.reject(context.error);
          return Promise.reject(
            (context.error = new RequestError(
              error.status ?? 500,
              error.message || "服务错误"
            ))
          );
        })
        .finally(() => {
          tasks.delete(seed);
          if (tasks.size === 0) {
            context.loading = false;
            context.pending = false;
          }
        });
    },
  });
  return context;
};
