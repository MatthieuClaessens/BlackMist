// @ts-ignore
import { StartTor, StopTor, CheckIP } from "../../wailsjs/go/main/App";

export const TorService = {
    async connect(): Promise<string> {
        try {
            return await StartTor();
        } catch (err) {
            throw err;
        }
    },

    async disconnect(): Promise<string> {
        try {
            return await StopTor();
        } catch (err) {
            return "Error during disconnection";
        }
    },

    async getIP(): Promise<string> {
        try {
            return await CheckIP();
        } catch (err) {
            return "Connecting...";
        }
    }
};