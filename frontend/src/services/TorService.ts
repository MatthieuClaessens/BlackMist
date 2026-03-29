/* eslint-disable */
import * as App from "../../wailsjs/go/main/App";

export const TorService = {
    async connect(): Promise<string> {
        return await App.StartTor();
    },

    async disconnect(): Promise<string> {
        try {
            return await App.StopTor();
        } catch {
            return "Error during disconnection";
        }
    },

    async getIP(): Promise<string> {
        try {
            return await App.CheckIP();
        } catch {
            return "Connecting...";
        }
    },

    async getPing(): Promise<number> {
        try {
            const p = await App.GetPing();
            return typeof p === 'number' ? p : 0;
        } catch {
            return 0;
        }
    }
};