import styles from "./styles.module.css";
import { Link, useNavigate } from "react-router-dom";
import HandleError from "../../components/error";
import React from "react";
import { HandleButton } from "../../components/button";
import { HandleRegisterOrganizationForm } from "./components/form";
import { HandleRegisterOrganizationTitle } from "./components/title";
import registerOrganizationServer from "../../server/register-organization";
export const PathRegisterOrganization = "/register-organization";

export function RegisterOrganization() {
  const router = useNavigate();
  const [error, setError] = React.useState<string | null>("");
  const [name, setName] = React.useState<string | null>("");
  const [document, setDocument] = React.useState<string | null>("");
  const [password, setPassword] = React.useState<string | null>("");
  const [logo, setLogo] = React.useState<string | null>("");

  const register = async () => {
    const createdOrganization = await registerOrganizationServer.Run({
      name,
      document,
      password,
      logo,
    });

    if (createdOrganization.error) {
      return setError(createdOrganization.message);
    }

    return router("/login-organization");
  };

  return (
    <section className={styles.container}>
      <HandleError error={error} function={setError} />
      <div className={styles.menu}>
        <HandleRegisterOrganizationTitle />
        <HandleRegisterOrganizationForm
          name={setName}
          document={setDocument}
          password={setPassword}
          logo={setLogo}
        />
        <HandleButton onClick={async () => await register()}>
          Register
        </HandleButton>
        <Link to={"/login"}>back</Link>
      </div>
    </section>
  );
}
