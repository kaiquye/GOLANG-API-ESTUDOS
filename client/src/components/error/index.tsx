import React from "react";
import styles from "./styles.module.css";
interface IProps {
  error: string | null;
  function: (error: string | null) => void;
}

export default function HandleError(error: IProps) {
  const [open, setOpen] = React.useState<string>();

  if (error.error != null) {
    setTimeout(() => {
      setOpen("none");
      error.function(null);
      setTimeout(() => {
        setOpen("flex");
      }, 10);
    }, 1650);
  }

  return (
    <>
      {error.error && (
        <>
          <div className={styles.container} style={{ display: open }}>
            <h1>{error.error}</h1>
          </div>
        </>
      )}
    </>
  );
}
