import React from "react";
import { useNavigate } from "react-router-dom";
import styles from "./styles.module.css";
import HandleError from "../../components/error";
import { HandleRegisterVacancyForm } from "./components/form";
import { HandleButton } from "../../components/button";
import { Link } from "react-router-dom";
import RegisterVacancyServer from "../../server/register-vacancy";
export const PathRegisterVacancy = "/register-vacancy";

export function RegisterVacancy() {
  const router = useNavigate();
  const [error, setError] = React.useState<string | null>("");

  const [title, setTitle] = React.useState<string | null>("");
  const [body, setBody] = React.useState<string | null>("");
  const [location, setLocation] = React.useState<number | null>("");
  const [level, setLevel] = React.useState<number | null>("");
  const [salary_range, setSalaryRange] = React.useState<number | null>("");
  const [organization_id, setOrganizationId] = React.useState<string | null>(
    ""
  );

  const Register = async () => {
    const created = await RegisterVacancyServer.Run({
      tilte: title,
      body,
      location,
      level,
      salary_range,
      organization_id,
    });
    if (created.error) {
      return setError(created.message);
    }

    return router("/list-vacancys");
  };

  return (
    <section className={styles.container}>
      <HandleError error={error} function={setError} />
      <div className={styles.menu}>
        <div className={styles.title}>
          <h1>Register new vacancy</h1>
        </div>
        <HandleRegisterVacancyForm
          body={setBody}
          level={setLevel}
          organization_id={setOrganizationId}
          salary_range={setSalaryRange}
          location={setLocation}
          tilte={setTitle}
        />
        <HandleButton onClick={async () => await Register()}>
          Login
        </HandleButton>
        <Link to={"/list-vacancy"}>List all vacancy</Link>
      </div>
    </section>
  );
}
