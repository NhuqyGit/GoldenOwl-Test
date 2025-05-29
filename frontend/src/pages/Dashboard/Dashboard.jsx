import "./dashboard.scss";
import { useStudentScores } from "../../contexts/StudentScoresProvider";
import Top10 from "../../components/Top10/Top10";
import Loader from "../../components/Loader/Loader";

const Dashboard = () => {
    const { top10GroupA, loading } = useStudentScores();
    useStudentScores;
    return (
        <div className="dashboard-container">
            <h2>Top 10 - Group A</h2>
            {loading && <Loader />}
            {!loading && <Top10 data={top10GroupA} />}
        </div>
    );
};

export default Dashboard;
