import { HandleInput } from "../../../../components/input";
import styles from "./styles.module.css";

interface IProps {
  name: (name: string) => void;
  document: (name: string) => void;
  password: (name: string) => void;
  logo: (name: string) => void;
}
export function HandleRegisterOrganizationForm(props: IProps) {
  return (
    <form className={styles.menu}>
      <label>name: </label>
      <HandleInput
        placeholder={"organization name"}
        onChange={(event) => props.name(event.target.value)}
      />
      <label>document: </label>
      <HandleInput
        placeholder={"organization document"}
        onChange={(event) => props.document(event.target.value)}
      />
      <label>password: </label>
      <HandleInput
        placeholder={"organization password"}
        onChange={(event) => props.password(event.target.value)}
      />
      <label>logo: </label>
      <HandleInput
        placeholder={"organization logo"}
        onChange={(event) => props.logo(event.target.value)}
      />
    </form>
  );
}
