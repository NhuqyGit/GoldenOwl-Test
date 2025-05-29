import React from "react";
import { ResponsiveBar } from "@nivo/bar";
import BarChartToolTip from "./BarChartToolTip";
import "./Chart.scss";

const rawData = {
    dia_li: { Average: 96226, Excellent: 218515, Good: 382087, Poor: 7854 },
    gdcd: { Average: 16886, Excellent: 384222, Good: 181440, Poor: 1061 },
    hoa_hoc: { Average: 88447, Excellent: 93333, Good: 144959, Poor: 19779 },
    lich_su: { Average: 200392, Excellent: 138533, Good: 342577, Poor: 24712 },
    ngoai_ngu: {
        Average: 363532,
        Excellent: 133483,
        Good: 219652,
        Poor: 196038,
    },
    ngu_van: { Average: 141056, Excellent: 377879, Good: 513116, Poor: 18050 },
    sinh_hoc: { Average: 116263, Excellent: 34438, Good: 182049, Poor: 9628 },
    toan: { Average: 258654, Excellent: 198392, Good: 505836, Poor: 82731 },
    vat_li: { Average: 79272, Excellent: 94146, Good: 148641, Poor: 23556 },
};

// Chuyển đổi dữ liệu sang định dạng mà Nivo chấp nhận
const subjectMap = {
    toan: "Math",
    ngu_van: "Literature",
    ngoai_ngu: "English",
    vat_li: "Physics",
    hoa_hoc: "Chemistry",
    sinh_hoc: "Biology",
    lich_su: "History",
    dia_li: "Geography",
    gdcd: "Civic Education",
};

const chartData = Object.entries(rawData).map(([subject, scores]) => ({
    subject: subjectMap[subject] || subject, // fallback nếu có subject lạ
    ...scores,
}));

const NivoGroupedBarChart = ({ reportData }) => {
    return (
        <div className="nivoGroupedBarChart">
            <h3>Distribution of Student Performance by Subject</h3>
            <ResponsiveBar
                data={reportData}
                keys={["Excellent", "Good", "Average", "Poor"]}
                indexBy="subject"
                margin={{ top: 40, right: 130, bottom: 60, left: 80 }}
                padding={0.3}
                groupMode="grouped"
                colors={{ scheme: "nivo" }}
                borderColor={{ from: "color", modifiers: [["darker", 1.6]] }}
                axisTop={null}
                axisRight={null}
                axisBottom={{
                    tickSize: 5,
                    tickPadding: 5,
                    tickRotation: -30,
                    legend: "Subjects",
                    legendPosition: "middle",
                    legendOffset: 50,
                }}
                axisLeft={{
                    tickSize: 5,
                    tickPadding: 5,
                    tickRotation: 0,
                    legend: "Number of students",
                    legendPosition: "middle",
                    legendOffset: -60,
                }}
                labelSkipWidth={12}
                labelSkipHeight={12}
                labelTextColor={{ from: "color", modifiers: [["darker", 1.6]] }}
                legends={[
                    {
                        dataFrom: "keys",
                        anchor: "bottom-right",
                        direction: "column",
                        justify: false,
                        translateX: 120,
                        translateY: 0,
                        itemsSpacing: 2,
                        itemWidth: 100,
                        itemHeight: 20,
                        itemDirection: "left-to-right",
                        itemOpacity: 0.85,
                        symbolSize: 20,
                    },
                ]}
                role="application"
                ariaLabel="Nivo bar chart"
                tooltip={BarChartToolTip}
            />
        </div>
    );
};

export default NivoGroupedBarChart;
