import "./Chart.scss";

const BarChartToolTip = ({ id, value, color, indexValue }) => {
    return (
        <div
            className="tooltip"
            style={{
                border: `1.5px solid ${color}`,
            }}
        >
            <div className="key-detail">
                <div>
                    <strong>Level:</strong>
                    {id}
                </div>
                <div>
                    <strong>{indexValue}:</strong>
                    {value.toLocaleString()}
                </div>
            </div>
        </div>
    );
};

export default BarChartToolTip;
