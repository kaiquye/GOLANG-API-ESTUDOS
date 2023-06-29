import { HandleInput } from "../../../../components/input";
import styles from "./styles.module.css";
import ListOrganizationsServer from "../../../../server/list-organizations";
import React from "react";
import { IOrganization } from "../../../../types/organization-types";

interface IProps {
  tilte: (target: string) => void;
  body: (target: string) => void;
  location: (target: number) => void;
  level: (target: number) => void;
  salary_range: (target: number) => void;
  organization_id: (target: string) => void;
}
export function HandleRegisterVacancyForm(props: IProps) {
  const [organizations, setOrganizations] = React.useState<IOrganization[]>([]);

  const listOrganizations = async () => {
    const response = await ListOrganizationsServer.Run();
    setOrganizations(response);
  };

  React.useEffect(() => {
    listOrganizations();
  }, []);

  return (
    <form className={styles.menu}>
      <label>document: </label>
      <HandleInput
        placeholder={"organization document"}
        onChange={(event) => props.tilte(event.target.value)}
      />
      <label>body: </label>
      <HandleInput
        placeholder={"organization body"}
        onChange={(event) => props.body(event.target.value)}
      />
      <label>level: </label>
      <HandleInput
        placeholder={"organization level"}
        onChange={(event) => props.level(event.target.value)}
      />
      <label>Location: </label>
      <HandleInput
        placeholder={"organization location"}
        onChange={(event) => props.location(event.target.value)}
      />
      <label>Salary range: </label>
      <HandleInput
        placeholder={"organization salary range"}
        onChange={(event) => props.salary_range(event.target.value)}
      />
      <label>Organization: </label>
      <select onChange={(event) => props.organization_id(event.target.value)}>
        <option defaultChecked={true}>select organization</option>
        {organizations.map((organization) => (
          <option value={organization.Org_id} key={organization.Org_id}>
            {organization.name}
          </option>
        ))}
      </select>
    </form>
  );
}
