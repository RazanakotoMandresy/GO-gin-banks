import AxiosInstances from "../AxiosInstances";

export const RegisterFunc = async (registerUser) => {
  return await AxiosInstances.post("/user/register", registerUser);
};
export const LoginFunc = async (logUser) => {
  return await AxiosInstances.post("/user/login", logUser);
};
export const GetUser = async (authentified) => {
  return await AxiosInstances.get("/user/logedUser", authentified);
};
