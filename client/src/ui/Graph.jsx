
const Chart = require('chart.js');
const PropTypes = require('prop-types');
const React = require('react');

class Graph extends React.Component {
  static propTypes = {
    data: PropTypes.array.isRequired
  };

  constructor(props) {
    super(props);
    this.chart = null;
  };

  componentDidMount() {
    setTimeout(() => {
      if (!this.chart) {
        this.drawGraph();
      }
    }, 5000);
  };

  componentDidUpdate(prevProps) {
    this.drawGraph();
  };

  componentWillUnmount() {
    if (this.chart) {
      this.chart.destroy();
    }
  };

  drawGraph = () => {
    const config = {
      type: 'doughnut',
      data: {
        labels: ['Recv.', 'Sent'],
        datasets: [{
          label: 'Bytes',
          data: this.props.data,
          backgroundColor: ['#5596e6', '#3d70b2'],
          borderColor: ['#4586d6', '#2d60a2'],
          borderWidth: 1
        }]
      },
      options: {}
    };
    const ctx = document.getElementById('doughnut-chart').getContext('2d');

    this.chart = new Chart(ctx, config);
  };

  render() {
    return (
      <canvas id="doughnut-chart" />
    );
  };
}

module.exports = Graph;
