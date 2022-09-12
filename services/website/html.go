package website

import (
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
}

func parseIndexTemplate() (*template.Template, error) {
	return template.New("index").Parse(`
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

    <link rel="stylesheet" href="https://storage.googleapis.com/eden_brand/img/pure-min.css" integrity="sha384-yHIFVG6ClnONEA5yB5DJXfW2/KC173DIQrYoZMEtBvGzmf0PKiGyNEqe9N6BNDBH" crossorigin="anonymous">
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

        .box { background:var(--darkblue);border-radius:10px;padding:15px 30px;width:100%;height:100%; }
        .config-table { width:100%; }

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
                        <a class="button" target="_block" href="https://v2.docs.edennetwork.io">Relay Docs</a>
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
                <table class="pure-table pure-table-horizontal pure-table-striped">
                    <thead>
                        <tr>
                            <th>Epoch</th>
                            <th>Slot</th>
                            <th>Block number</th>
                            <!--<th>Parent hash</th>-->
                            <th>Block hash</th>
                            <th>Num tx</th>
                            <th>Value</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .Payloads }}
                        <tr>
                            <td>{{.Epoch}}</td>
                            <td><a href="/relay/v1/data/bidtraces/proposer_payload_delivered?slot={{.Slot}}">{{.Slot}}</a>
                            </td>
                            <td>{{.BlockNumber}}</td>
                            <td>{{.BlockHash}}</td>
                            <!--<td>{{.ParentHash}}</td>-->
                            <td>{{.NumTx}}</td>
                            <td>{{.Value}}</td>
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
            <a class="button" target="_blank" href="/relay/v1/data/bidtraces/proposer_payload_delivered?limit=10">Data API</a> <a class="button" target="_blank" href="https://v2.docs.edennetwork.io">Docs</a></p>
        </div>
    </div>
</body>
</html>
`)
}
