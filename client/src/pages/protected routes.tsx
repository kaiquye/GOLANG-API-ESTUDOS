import { Navigate, Outlet } from "react-router-dom";

interface IProtectedParams {
  accessToken: string | null;
  redirectPath: "/login-organization";
}
export function ProtectedRoutes(input: IProtectedParams) {
  // if (!input?.accessToken) {
  //   return <Navigate to={input.redirectPath} replace />;
  // }

  return <Outlet />;
}
