package website

import (
	"math/big"
	"text/template"

	"github.com/flashbots/mev-boost-relay/database"
)

type StatusHTMLData struct {
	Network                     string
	RelayPubkey                 string
	ValidatorsTotal             string
	ValidatorsRegistered        string
	BellatrixForkVersion        string
	GenesisForkVersion          string
	GenesisValidatorsRoot       string
	BuilderSigningDomain        string
	BeaconProposerSigningDomain string
	HeadSlot                    string
	NumPayloadsDelivered        string
	Payloads                    []*database.DeliveredPayloadEntry
	ValueLink                   string
	ValueOrderIcon              string
	ShowConfigDetails           bool
	LinkBeaconchain             string
	LinkEtherscan               string
}

func weiToEth(wei string) string {
	weiBigInt := new(big.Int)
	weiBigInt.SetString(wei, 10)
	ethValue := weiBigIntToEthBigFloat(weiBigInt)
	return ethValue.String()
}

func weiBigIntToEthBigFloat(wei *big.Int) (ethValue *big.Float) {
	// wei / 10^18
	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue = new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	return
}

var funcMap = template.FuncMap{
	"weiToEth": weiToEth,
}

func ParseIndexTemplate() (*template.Template, error) {
	return template.New("index").Funcs(funcMap).Parse(`
<!DOCTYPE html>
<html lang="en" class="no-js">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">

    <meta name="viewport" content="width=device-width,initial-scale=1;">

    <title>Eden Relay - {{ .Network }}</title>

    <meta name="description"
        content="Eden testing relay for maximal extractable value in Ethereum proof-of-stake.">
    <link data-react-helmet="true" rel="shortcut icon" href="https://storage.googleapis.com/eden_brand/img/favicon-32x32.png">

    <link rel="stylesheet" data-href="https://fonts.googleapis.com/css2?family=Titillium+Web:wght@300;400;600;700&amp;display=swap"/><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />

    <link rel="stylesheet" href="https://cdn.rawgit.com/Arch-Matt/M-grid/ff2ee025/m-grid-min.1.0.0.css">


    <style type="text/css">
        :root{
            --green: #caff00;
            --lightblue:#2b305a;
            --darkblue:#171c47;
        }
        body {
            padding: 10px 15px;
            background: var(--lightblue);
            color: #fff;
            font-family:'Titillium Web', serif;
            margin: 0;
        }

        pre {
            text-align: left;
        }

        hr {
            border-top: 1px solid #e5e5e5;
            margin: 40px 0;
        }

        tt {
            font-size: 1.2em;
            background:var(--lightblue);
        }

        li {
            margin: 2px 0px;
        }

        .pure-table thead {
            background-color:var(--lightblue);
        }

        .pure-table tr:hover td {
            background:var(--lightblue) !important;
        }

        .pure-table { width:100%;min-width:600px; }

        

        .table-holder { width:100%;overflow-x: auto; }

        .text-green { color:var(--green); }

        .m-grid { margin:0; }


        ul { list-style:none;padding:0; }

        .stat-block { margin: 0;list-style: none;display:inline-table;table-layout: fixed;padding:0;background:var(--green);color:var(--darkblue);font-weight:400;border-radius:10px;font-size:1.125rem; }
        .stat-block li { display:table-cell;padding:15px 30px;border-right:2px solid var(--lightblue); }
        .stat-block li:last-child { border:none; }
        .stat-block span { display:block;margin:0 auto;font-weight:700;font-size:3rem; }

        input[type="text"] { border:none;background:var(--lightblue);padding:5px 10px;color:#fff;width:100%;text-align:center;margin-top:-10px;display:block; }

        .box { background:var(--darkblue);border-radius:10px;padding:15px 30px 30px 30px;width:100%;height:100%; }
        .config-table { width:100%; }

        a:link    { color:var(--green); text-decoration: none; }
        a:visited { color:var(--green); text-decoration: none; }
        a:hover   { color: #fff; text-decoration: none; }
        a:active  { color: #fff; text-decoration: none; }

        a.button { background:var(--green);color:var(--darkblue);display:inline-block;margin:7px 7px 7px 7px;padding:15px 30px;width:auto;border-radius:5px;text-decoration: none;font-weight:700; }
        a.button:hover { background:#fff; }

        h1 { font-size:3rem; }
        h2 { margin-top:0; }
        h3 { font-size:1.125rem; }

        .logobar { width:calc(100% + 30px);margin-top: -10px;margin-left: -15px;background:var(--darkblue);padding:15px }
        .logo { height:30px; }

        @media(max-width: 768px){
            .stat-block, .stat-block li { display:block;width:100%; }
            .stat-block li { border:none;border-bottom:2px solid var(--darkblue); }
            .stat-block li:last-child { border:none; }
            input[type="text"] { text-align:left; }
        }

		.icons-container {
            display: flex;
            flex-direction: row;
            justify-content: center;
            align-items: center;
            max-width: 50px;
        }

        .icons-container a {
			text-align: center;
            display:block;
            width:18px;
            height:18px;
        }

		.img-beaconchain {
            background-repeat: no-repeat;
            background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='263.27 168.54 81.35 101.17'%3E%3Cpath d='M 341.287 198.092 L 340.508 198.289 C 338.293 198.85 337.797 198.496 336.846 196.409 C 335.8 193.738 334.475 191.177 332.89 188.769 C 325.505 178.925 315.692 174.652 303.269 176.552 C 301.781 176.788 300.881 176.336 300.517 174.869 C 300.294 173.884 300.041 172.998 299.829 172.053 C 299.445 170.36 300.001 169.513 301.711 169.208 C 312.434 167.239 322.146 169.631 330.887 175.962 C 337.423 180.697 341.702 187.076 344.363 194.666 C 344.98 196.33 344.484 197.216 342.704 197.708 L 342.704 197.738 Z M 329.865 201.006 C 328.166 201.439 327.842 199.894 326.578 196.94 C 323.522 189.645 315.884 185.189 307.831 186.003 C 305.252 186.259 303.461 186.712 303.117 185.343 C 302.055 181.041 302.237 180.795 306.688 180.421 C 317.858 179.693 328.237 186.055 332.385 196.172 C 332.476 196.419 332.567 196.665 332.678 196.911 C 333.73 199.342 333.366 200.081 330.796 200.77 Z M 323.866 199.923 C 323.927 200.071 323.988 200.228 324.058 200.376 C 324.726 201.882 324.483 202.345 322.824 202.778 L 322.227 202.926 C 321.135 203.201 320.963 202.237 320.133 200.396 C 318.102 195.813 313.219 193.07 308.125 193.652 C 306.455 193.829 305.302 194.125 305.09 193.268 C 304.422 190.581 304.543 190.423 307.406 190.167 C 314.489 189.592 321.142 193.536 323.866 199.923 Z M 274.143 222.713 C 271.866 216.439 271.442 209.669 272.919 203.172 C 273.778 199.234 275.144 198.673 278.473 200.839 C 285.554 205.446 292.515 209.935 300.082 214.838 C 300.82 213.528 301.802 212.278 302.459 210.851 C 302.775 210.139 302.883 209.355 302.773 208.586 C 302.267 205.525 303.785 202.827 306.648 202.138 C 309.515 201.493 312.366 203.272 312.96 206.076 C 313.579 208.883 311.831 211.669 308.985 212.416 C 308.512 212.509 308.076 212.73 307.73 213.056 C 307.558 213.312 304.624 217.899 304.624 217.899 L 315.014 224.889 C 318.525 227.242 322.005 229.575 325.546 231.879 C 329.086 234.182 328.975 235.639 325.738 238.179 C 308.954 251.233 282.104 244.972 274.143 222.713 Z M 263.271 269.704 L 263.271 269.655 C 263.758 268.742 276.643 244.586 277.269 243.948 C 278.645 242.55 298.322 252.808 299.525 254.019 C 299.778 254.275 304.988 269.682 304.988 269.682 L 304.988 269.704 Z'%3E%3C/path%3E%3C/svg%3E");
        }
    </style>
</head>

<body>

    <div class="logobar">
        <div class="m-container-1600">
            <a href="https://edennetwork.io"><img class="logo" src="https://storage.googleapis.com/eden_brand/img/logo.svg"></a>
        </div>
    </div>


    <div class="m-container-1600 text-center">
        <h1>
            Eden Relay - <span class="text-green">{{ .Network }}</span>
        </h1>



        <div class="m-section-15">
            <div class="m-grid m-grid-gap-30">
                <div class="m-xl-12 m-lg-7 m-md-12 m-sm-12">
                    <div class="box">
                         <h3>
                            Configuration
                        </h3>

                        <p>Relay Pubkey:</p>
                        <input type="text" value="{{ .RelayPubkey }}">

                        <p>Bellatrix fork version:</p>
                        <input type="text" value="{{ .BellatrixForkVersion }}">

                        <p>Genesis fork version:</p>
                        <input type="text" value="{{ .GenesisForkVersion }}">

                        <p>Genesis validators root:</p>
                        <input type="text" value="{{ .GenesisValidatorsRoot }}">
                        
                        <p>Builder signing domain:</p>
                        <input type="text" value="{{ .BuilderSigningDomain }}">

                        <p>Beacon proposer signing domain:</p>
                        <input type="text" value="{{ .BeaconProposerSigningDomain }}">
                    </div>
                </div>                
            </div>
        </div>           


        <div class="m-section-30">
            <h2>
                Stats
            </h2>

            <ul class="stat-block">
                <li><span>{{ .ValidatorsTotal }}</span>Validators total</li>
                <li><span>{{ .ValidatorsRegistered }}</span>Validators registered</li>
                <li><span>{{ .HeadSlot }}</span>Latest slot</li>
            </ul>
        </div>




        <div class="m-section-15">
            <div class="box">
                <h2>
                    Recently Delivered Payloads
                </h2>

                <div class="table-holder">
					<table class="pure-table pure-table-horizontal pure-table-striped" style="width:100%;">
						<thead>
							<tr>
								<th>Epoch</th>
								<th>Slot</th>
								<th>Block number</th>
								<th>
									Value (ETH{{.ValueOrderIcon}})
									<a href="{{.ValueLink}}">
										<svg id="icon-sort-default" style="float:right; width:16px;" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
											<path stroke-linecap="round" stroke-linejoin="round" d="M3 7.5L7.5 3m0 0L12 7.5M7.5 3v13.5m13.5 0L16.5 21m0 0L12 16.5m4.5 4.5V7.5" />
										</svg>
									</a>
								</th>
								<th>Num tx</th>
								<th>Block hash</th>
								<th></th>
							</tr>
						</thead>
						<tbody>
							{{$linkBeaconchain := .LinkBeaconchain}}
							{{$linkEtherscan := .LinkEtherscan}}
							{{ range .Payloads }}
							<tr>
								<td>{{.Epoch}}</td>
								<td>
									<a href="/relay/v1/data/bidtraces/proposer_payload_delivered?slot={{.Slot}}">{{.Slot}}</a>
								</td>
								<td>{{.BlockNumber}}</td>
								<td>{{.Value | weiToEth}}</td>
								<td>{{.NumTx}}</td>
								<td>{{.BlockHash}}</td>
								<td>
									<div class="icons-container">
										{{ if ne $linkBeaconchain "" }}
											<a class="img-beaconchain" href="{{$linkBeaconchain}}/block/{{.BlockNumber}}" target="_blank" alt="View block in beaconcha.in" title="View block in beaconcha.in"></a>
											&nbsp;
										{{ end }}
										{{ if ne $linkEtherscan "" }}
											<a href="{{$linkEtherscan}}/block/{{.BlockNumber}}" target="_blank"><img src="https://etherscan.io/images/favicon3.ico" width="18" height="18" alt="View block on Etherscan" title="View block on Etherscan" /></a>
										{{ end }}
									</div>
								</td>
							</tr>
							{{ end }}
						</tbody>
					</table>
                </div>
            </div>
        </div>



        <div class="m-section-15">
            <h2>
                {{.NumPayloadsDelivered}} payloads delivered</p>
            </h2>
            <a class="button" target="_blank" href="/relay/v1/data/bidtraces/proposer_payload_delivered?limit=10">Data API</a> <a class="button" target="_blank" href="https://docs.edennetwork.io">Docs</a></p>
        </div>
    </div>
</body>
</html>
`)
}
