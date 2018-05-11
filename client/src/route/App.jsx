
import '../style/custom.scss';

import fetch from '../lib/fetch';
import React from 'react';

import Footer from '../comp/Footer';
import Header from '../comp/Header';

// Cards
import Bandwidth from '../card/Bandwidth';
import BandwidthChart from '../card/BandwidthChart';
import Donate from '../card/Donate';
import Fee from '../card/Fee';
import MemPool from '../card/MemPool';
import Network from '../card/Network';
import Node from '../card/Node';
import NodeDetail from '../card/NodeDetail';
import Status from '../card/Status';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.interval = null;
    this.state = {
      blocks: 12345,
      connections: 125,
      country: 'United States',
      difficulty: 100.0,
      donationAddress: 'DONATION',
      ip: '192.168.0.1',
      incomingData: 1000.0,
      maxFee: 0.0,
      midFee: 0.0,
      minFee: 0.0,
      maxMemory: 0.0,
      network: 'nonet',
      networkHashPS: 0.0,
      outgoingData: 500.0,
      protocol: 77777,
      rank: 0,
      status: 'Offline',
      stakingStatus: 'Unknown',
      subversion: '1.1.1.1',
      totalData: 1500.0,
      transactions: 3,
      usedMemory: 340.0,
      version: 1
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
      <div className="bx--grid">
        <Header />
        <article>
          <div className="bx--row" style={{ padding: 20 }}>
            <MemPool
              max={ this.state.maxMemory }
              size={ this.state.usedMemory }
              transactions={ this.state.transactions } />
            <NodeDetail
              blocks={ this.state.blocks }
              connections={ this.state.connections }
              network={ this.state.network }
              status={ this.state.status } />
            <Node
              ip={ this.state.ip }
              protocol={ this.state.protocol }
              subversion={ this.state.subversion }
              version={ this.state.version } />
            <BandwidthChart
              incomingData={ this.state.incomingData }
              outgoingData={ this.state.outgoingData }
              totalData={ this.state.totalData } />
            <Network
              blocks={ this.state.blocks }
              difficulty={ this.state.difficulty }
              networkHashPS={ this.state.networkHashPS } />
            <Bandwidth
              incomingData={ this.state.incomingData }
              outgoingData={ this.state.outgoingData }
              totalData={ this.state.totalData } />
            <Fee
              max={ this.state.maxFee }
              mid={ this.state.midFee }
              min={ this.state.minFee } />
            <Donate
              donationAddress={ this.state.donationAddress } />
            <Status
              stakingStatus={ this.state.stakingStatus }
              status={ this.state.status } />
          </div>
        </article>
        <Footer />
      </div>
    );
  }
}

export default App;
