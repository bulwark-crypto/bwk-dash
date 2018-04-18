
const fetch = require('../lib/fetch');
const React = require('react');

const Footer = require('../component/Footer');
const Header = require('../component/Header');

// Cards
const Bandwidth = require('../component/Bandwidth');
const BandwidthChart = require('../component/BandwidthChart');
const Donate = require('../component/Donate');
const Fee = require('../component/Fee');
const MemPool = require('../component/MemPool');
const Network = require('../component/Network');
const Node = require('../component/Node');
const NodeDetail = require('../component/NodeDetail');
const Status = require('../component/Status');

class App extends React.Component {
  constructor(props) {
    super(props);

    this.interval = null;
    this.state = {
      blocks: 0,
      connections: 0,
      country: '',
      difficulty: 0.0,
      donationAddress: 'DONATION',
      ip: '',
      incomingData: 0.0,
      maxFee: 0.0,
      midFee: 0.0,
      minFee: 0.0,
      maxMemory: 0.0,
      network: 'nonet',
      networkHashPS: 0.0,
      outgoingData: 0.0,
      protocol: 0,
      rank: 0,
      status: 'Down',
      stakingStatus: 'Unknown',
      subversion: '',
      totalData: 0.0,
      transactions: 0,
      usedMemory: 0.0,
      version: 0
    };
  };

  componentDidCatch(error, info) {
    this.setState({ error: error.message });
  };

  componentDidMount() {
    fetch.getInfo()
      .then(info => this.setState({ ...info }))
      .catch(error => console.log(error));
  };

  render() {
    return (
      <div className="app">
        <Header />
        <div className="content">
          { !!this.state.error &&
            <div className="error">{ this.state.error }</div>
          }
          { !this.state.error &&
            <div className="row">
              <Node
                ip={ this.state.ip }
                protocol={ this.state.protocol }
                subversion={ this.state.subversion }
                version={ this.state.version } />
              <NodeDetail
                blocks={ this.state.blocks }
                connections={ this.state.connections }
                network={ this.state.network } />
              <MemPool
                size={ this.state.usedMemory }
                transactions={ this.state.transactions } />
            <div className="row">
            </div>
              <Bandwidth
                incomingData={ this.state.incomingData }
                outgoingData={ this.state.outgoingData }
                totalData={ this.state.totalData } />
              <Network
                blocks={ this.state.blocks }
                difficulty={ this.state.difficulty }
                networkHashPS={ this.state.networkHashPS } />
              <BandwidthChart
                incomingData={ this.state.incomingData }
                outgoingData={ this.state.outgoingData }
                totalData={ this.state.totalData } />
            <div className="row">
            </div>
              <Status
                stakingStatus={ this.state.stakingStatus }
                status={ this.state.status } />
              <Donate
                donationAddress={ this.state.donationAddress } />
              <Fee
                max={ this.state.maxFee }
                mid={ this.state.midFee }
                min={ this.state.minFee } />
            </div>
          }
        </div>
        <Footer donation={ this.state.donationAddress } />
      </div>
    );
  };
}

module.exports = App;
