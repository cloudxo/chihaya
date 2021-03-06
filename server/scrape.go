/*
 * This file is part of Chihaya.
 *
 * Chihaya is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Chihaya is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Chihaya.  If not, see <http://www.gnu.org/licenses/>.
 */

package server

import (
	"chihaya/config"
	"chihaya/database"
	cdb "chihaya/database/types"
	"chihaya/server/params"
	"io"

	"github.com/zeebo/bencode"
)

var scrapeInterval int

func init() {
	intervals := config.Section("intervals")
	scrapeInterval, _ = intervals.GetInt("scrape", 900)
}

func writeScrapeInfo(torrent *cdb.Torrent) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["complete"] = len(torrent.Seeders)
	ret["downloaded"] = torrent.Snatched
	ret["incomplete"] = len(torrent.Leechers)

	return ret
}

func scrape(qs string, user *cdb.User, db *database.Database, buf io.Writer) {
	qp, err := params.ParseQuery(qs)
	if err != nil {
		panic(err)
	}

	scrapeData := make(map[string]interface{})
	fileData := make(map[string]interface{})

	if qp.InfoHashes() != nil {
		db.TorrentsMutex.RLock()

		for _, infoHash := range qp.InfoHashes() {
			torrent, exists := db.Torrents[infoHash]
			if exists {
				if !isDisabledDownload(db, user, torrent) {
					fileData[infoHash] = writeScrapeInfo(torrent)
				}
			}
		}

		db.TorrentsMutex.RUnlock()
	} else {
		scrapeData["failure reason"] = "Scrape without info_hash is not supported"
	}

	scrapeData["files"] = fileData
	scrapeData["flags"] = map[string]interface{}{
		"min_request_interval": scrapeInterval,
	}
	// the following are for compatibility with clients that don't implement scrape flags
	scrapeData["interval"] = scrapeInterval
	scrapeData["min interval"] = scrapeInterval

	bufdata, err := bencode.EncodeBytes(scrapeData)
	if err != nil {
		panic(err)
	}

	_, err = buf.Write(bufdata)
	if err != nil {
		panic(err)
	}
}
