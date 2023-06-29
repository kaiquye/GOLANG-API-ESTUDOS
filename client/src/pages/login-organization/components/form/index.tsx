import { HandleInput } from "../../../../components/input";
import styles from "./styles.module.css";
import { HandleButton } from "../../../../components/button";

interface IProps {
  document: (name: string) => void;
  password: (name: string) => void;
}
export function HandleLoginOrganizationForm(props: IProps) {
  return (
    <form className={styles.menu}>
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
    </form>
  );
}
