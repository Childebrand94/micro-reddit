import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App"; // Assuming App.tsx is in the same directory
import "./index.css";
import { FilterProvider } from "./context/FiletProvider.tsx";

const root = ReactDOM.createRoot(document.getElementById("root")!);
root.render(
    <React.StrictMode>
        <FilterProvider>
            <App />
        </FilterProvider>
    </React.StrictMode>,
);
