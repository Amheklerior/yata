import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { QueryClientProvider } from "@tanstack/react-query";
import { queryClient } from "./lib/query.ts";
import { NotificationProvider } from "./contexts/NotificationCtxProvider.tsx";
import { NotificationCenter } from "./components/NotificationCenter.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <NotificationProvider>
        <App />
        <NotificationCenter />
        <div className="paper-bg" />
      </NotificationProvider>
    </QueryClientProvider>
  </StrictMode>,
);
