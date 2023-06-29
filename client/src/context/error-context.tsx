import React from "react";

interface IProps {
  error: string | null;
  function: (error: string | null) => void;
}

export const ErrorContext = React.createContext<IProps>({} as any);

export function ProviderErrorContext({ children }) {
  const [error, setError] = React.useState<string | null>(null);
  return (
    <ErrorContext.Provider value={{ error, function: setError }}>
      {children}
    </ErrorContext.Provider>
  );
}

export const useErrorContext = () => React.useContext(ErrorContext);
