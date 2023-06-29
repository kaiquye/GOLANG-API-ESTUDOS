import React from "react";
import { useNavigate } from "react-router-dom";
import styles from "./styles.module.css";
import HandleError from "../../components/error";
import { HandleLoginOrganizationForm } from "./components/form";
import { HandleButton } from "../../components/button";
import { Link } from "react-router-dom";
import LoginOrganizationServer from "../../server/login-organization";
import { AccessTokenStorage } from "../../server/storage/access-token-storage";

export const PathLoginOrganization = "/login-organization";

export function LoginOrganization() {
  const router = useNavigate();
  const [error, setError] = React.useState<string | null>("");
  const [document, setDocument] = React.useState<string | null>("");
  const [password, setPassword] = React.useState<string | null>("");

  const login = async () => {
    const logged = await LoginOrganizationServer.Run({
      document,
      password,
    });
    if (logged.error) {
      return setError("document or password invalid.");
    }
    console.log(logged);
    AccessTokenStorage.post(logged.accessToken);
    return router("/list-vacancy");
  };

  return (
    <section className={styles.container}>
      <HandleError error={error} function={setError} />
      <div className={styles.menu}>
        <div className={styles.title}>
          <h1>Login organization</h1>
        </div>
        <HandleLoginOrganizationForm
          document={setDocument}
          password={setPassword}
        />
        <HandleButton onClick={async () => await login()}>Login</HandleButton>
        <Link to={"/register-organization"}>Register your organization</Link>
      </div>
    </section>
  );
}
