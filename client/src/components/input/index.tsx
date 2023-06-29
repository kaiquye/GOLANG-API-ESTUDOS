import styles from "./styles.module.css";
interface IProps {
  onChange?: (value: any) => void;
  placeholder?: string;
}

export function HandleInput(props: IProps) {
  return (
    <input
      className={styles.HandleInput}
      onChange={props.onChange}
      placeholder={props.placeholder}
    />
  );
}
