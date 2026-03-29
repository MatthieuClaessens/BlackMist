import './App.css';
import TitleBar from './components/layout/TitleBar';
import HomePage from './pages/HomePage';

function App() {

    return (
        <div className="flex flex-col h-screen w-screen bg-[#060608] overflow-hidden select-none">
            <TitleBar />
            <main className="flex-1 flex items-center justify-center">
                <HomePage />
            </main>
        </div>
    );
}

export default App
