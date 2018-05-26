
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
      blocks: 0,
      connections: 0,
      country: 'United States',
      difficulty: 0.0,
      donationAddress: 'BURN',
      ip: '192.168.0.1',
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
      status: 'Offline',
      stakingStatus: 'Unknown',
      subversion: '0.0.0.0',
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
      .then(info => this.setState({ ...info }, () => {
        this.fetchData();
      }))
      .catch(error => console.log(error));
  };

  componentWillUnmount() {
    if (this.interval) {
      clearTimeout(this.interval);
      this.interval = null;
    }
  };

  fetchData = () => {
    if (this.interval) {
      clearTimeout(this.interval);
    }

    this.interval = setTimeout(() => {
      fetch.getInfo()
        .then(info => this.setState({ ...info }, () => {
          this.fetchData();
        }))
        .catch(error => console.log(error));
    }, 30 * 1000); // 30 secs
  };

  render() {
    return (
      <div>
        <Header />
        <div className="bx--grid">
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
      </div>
    );
  }
}

export default App;
