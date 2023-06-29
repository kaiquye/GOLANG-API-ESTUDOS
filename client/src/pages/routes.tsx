import { BrowserRouter, Route, Routes } from "react-router-dom";
import { LoginOrganization, PathLoginOrganization } from "./login-organization";
import { ProtectedRoutes } from "./protected routes";
import {
  PathRegisterOrganization,
  RegisterOrganization,
} from "./register-organization";
import { ProviderErrorContext } from "../context/error-context";
import { PathRegisterVacancy, RegisterVacancy } from "./register-vacancy";
import { ListVacancies, PathListVacancy } from "./list-vacancies";
import { HandleNavBar } from "../components/navBar";

export function AppRoutes() {
  const accessToken = localStorage.getItem("access_token");

  return (
    <BrowserRouter>
      <HandleNavBar />
      <ProviderErrorContext>
        <Routes>
          <Route path={PathLoginOrganization} element={<LoginOrganization />} />
          <Route
            element={
              <ProtectedRoutes
                accessToken={accessToken}
                redirectPath={"/login-organization"}
              />
            }
          >
            <Route
              path={PathRegisterOrganization}
              element={<RegisterOrganization />}
            />
            <Route path={PathRegisterVacancy} element={<RegisterVacancy />} />
            <Route path={PathListVacancy} element={<ListVacancies />} />
          </Route>
        </Routes>
      </ProviderErrorContext>
    </BrowserRouter>
  );
}
