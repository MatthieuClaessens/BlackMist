export { };

declare global {
    interface Window {
        go: {
            main: {
                App: {
                    CheckIP(): Promise<string>;
                    GetPing(): Promise<number>;
                    StartTor(): Promise<string>;
                    StopTor(): Promise<string>;
                };
            };
        };
    }
}