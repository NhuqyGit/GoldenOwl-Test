import "./App.css";
import MainContainer from "./components/Main/MainContainer";
import Menu from "./components/Menu/Menu";
import { StudentScoreProvider } from "./contexts/StudentScoresProvider";

function App() {
    return (
        <div className="app">
            <StudentScoreProvider>
                <Menu />
                <MainContainer />
            </StudentScoreProvider>
        </div>
    );
}

export default App;
