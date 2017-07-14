/*
 * Copyright (C) 2016-2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */
import React from 'react';
import Vault from '../models/vault';
import {T} from './Utils';


var Footer = React.createClass({
	getInitialState: function() {
      return {version: null};
  },

  componentDidMount: function() {
    this.getVersion();
  },

  getVersion: function() {
    var self = this;
    Vault.version().then(function(response) {
      var data = JSON.parse(response.body);
      self.setState({version: data.version});
    });
  },

	render: function() {
		return (
			<footer className="spacer">
				<div className="row">
					<p><small>{T('version')}: {this.state.version}</small></p>
				</div>
			</footer>
		);
	}
});

export default Footer;
