import NivoGroupedBarChart from "../../components/Chart/GroupChart";
import "./Report.scss";
import ContentLoader from "react-content-loader";
import { useStudentScores } from "../../contexts/StudentScoresProvider";
import Loader from "../../components/Loader/Loader";

const ChartSkeleton = () => (
    <ContentLoader
        speed={2}
        width="100%"
        height={500}
        backgroundColor="#f3f3f3"
        foregroundColor="#ecebeb"
    >
        <rect x="0" y="0" rx="5" ry="5" width="100%" height="500" />
    </ContentLoader>
);

const Report = () => {
    const { reportData, loadingReport } = useStudentScores();

    return (
        <div className="report-container">
            {/* {loading && <ChartSkeleton />} */}
            {loadingReport && <Loader />}
            {!loadingReport && reportData && (
                <NivoGroupedBarChart reportData={reportData} />
            )}
            {!loadingReport && !reportData && <div>Failed to load data</div>}
        </div>
    );
};

export default Report;
