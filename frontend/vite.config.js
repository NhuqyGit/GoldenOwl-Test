import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            "/student-scores": {
                target: import.meta.env.VITE_API_HOST,
                changeOrigin: true,
                secure: false,
            },
        },
    },
});
