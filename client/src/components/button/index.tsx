import style from "./style.module.css";

interface IProps {
  children?: any;
  onClick?: () => void;
}
export function HandleButton(props: IProps) {
  return (
    <button type={"submit"} onClick={props.onClick} className={style.button}>
      {props.children}
    </button>
  );
}
