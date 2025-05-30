import "./Dashboard.scss";
import { useStudentScores } from "../../contexts/StudentScoresProvider";
import Top10 from "../../components/Top10/Top10";
import Loader from "../../components/Loader/Loader";

const Dashboard = () => {
    const { top10GroupA, loadingTop10 } = useStudentScores();
    useStudentScores;
    return (
        <div className="dashboard-container">
            <h2>Top 10 - Group A</h2>
            {loadingTop10 && <Loader />}
            {!loadingTop10 && <Top10 data={top10GroupA} />}
        </div>
    );
};

export default Dashboard;
