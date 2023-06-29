import React from "react";
import styles from "./styles.module.css";
import { IVacancy } from "../../types/vacancy-types";
import ListVacancysServer from "../../server/list-vacancys";
import { HandleInput } from "../../components/input";
import { useNavigate } from "react-router-dom";
export const PathListVacancy = "/list-vacancy";

export function ListVacancies() {
  const modalRef = React.useRef<HTMLDivElement | null>(null);
  const router = useNavigate();
  const [modalStatus, setModalStatus] = React.useState<"open" | "closed">(
    "open"
  );
  const [vacancys, setVacancys] = React.useState<IVacancy[]>([]);

  const modalShow = {
    closed: () => {
      modalRef.current!.style.display = "none";
      setModalStatus(modalStatus == "closed" ? "open" : "closed");
    },
    open: () => {
      modalRef.current!.style.display = "flex";
      setModalStatus(modalStatus == "open" ? "closed" : "open");
    },
  };

  const getVacacys = async () => {
    const response = await ListVacancysServer.Run();
    setVacancys(response);
  };

  const toVacancies = () => {
    return router("/register-vacancy");
  };
  const toOrganization = () => {
    return router("/register-organization");
  };

  React.useEffect(() => {
    getVacacys();
  }, []);

  return (
    <section className={styles.container}>
      <div ref={modalRef} className={styles.Modal}>
        <h2>Send your data to the company</h2>
        <form>
          <label>Document:</label>
          <HandleInput placeholder={"your document"} />
          <label>Email:</label>
          <HandleInput placeholder={"your Email"} />
          <label>Level:</label>
          <HandleInput placeholder={"your Level"} />
          <label>Location:</label>
          <HandleInput placeholder={"your Location"} />
          <label>Resume:</label>
          <textarea />
          <button>Apply</button>
        </form>
      </div>
      <div className={styles.title}>
        <h2>Jobs</h2>
        <img
          src={`https://img.uxwing.com/wp-content/themes/uxwing/download/buildings-architecture-real-estate/job-icon.png`}
        />
        <button onClick={toVacancies} className={styles.BtnMenu}>
          Register your vacancy
        </button>
        <button onClick={toOrganization} className={styles.BtnMenu}>
          register your company
        </button>
      </div>
      <div className={styles.listVacancy}>
        <ul className={styles.list}>
          {vacancys.map((vacancy) => (
            <li>
              <h2>
                <label style={{ color: "rgba(196,192,192,0.62)" }}>
                  Title:
                </label>{" "}
                {vacancy.tilte}
              </h2>
              <div>
                <label
                  style={{
                    color: "#bdbdbd",
                    fontSize: "13px",
                    paddingRight: "4px",
                  }}
                >
                  Body:
                </label>{" "}
                {vacancy.body}
              </div>
              <div>
                <label
                  style={{
                    color: "#bdbdbd",
                    fontSize: "13px",
                    paddingRight: "4px",
                  }}
                >
                  Location:
                </label>
                {vacancy.location}
              </div>
              <div>
                <label
                  style={{
                    color: "#bdbdbd",
                    fontSize: "13px",
                    paddingRight: "4px",
                  }}
                >
                  Salary Range:
                </label>
                {vacancy.salary_range}
              </div>
              <div>
                <label
                  style={{
                    color: "#bdbdbd",
                    fontSize: "13px",
                    paddingRight: "4px",
                  }}
                >
                  Level:
                </label>
                {vacancy.level}
              </div>
              <button
                onClick={() => {
                  modalShow[modalStatus]();
                }}
              >
                Apply
              </button>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}
