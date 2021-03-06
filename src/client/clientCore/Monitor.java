package clientCore;

import java.io.IOException;

import javax.swing.JOptionPane;

import org.json.JSONException;

import packageManaging.Encoder;
import packageManaging.InitializeClientMessage;
import packageManaging.StageFlipper;

import clientHandlers.LobbyHandler;
import clientHandlers.SnakeHandler;
import clientHandlers.TiarHandler;
import clientNetworking.Connection;
import clientNetworking.NetManager;

public class Monitor {

	public LobbyHandler LH;
	public TiarHandler TH;
	public SnakeHandler SH;
	public Connection conn;
	public NetManager net;
	public String userName;
	public int stage;
	public StageFlipper lastMsg = null;

	public Monitor() {

	}


	public void init() {
	String serverConnectStr = "Enter server: ";
		while (true) {
			String hostStr = JOptionPane.showInputDialog(serverConnectStr, Connection.DEFAULT_HOST + ":" + String.valueOf(Connection.DEFAULT_PORT));
			if (hostStr == null) {
				return;
			}
			try {
				this.conn = new Connection(hostStr);
				break;
			} catch (NumberFormatException e) {
				serverConnectStr = "Invalid port. Enter server:";
			}
		}

		this.net = new NetManager(conn);
		try {
			net.connectToServer();
			userName = (String) JOptionPane.showInputDialog("Write username!");
			if (userName == null) {
				return;
			} else if (userName.length() == 0) {
				JOptionPane.showMessageDialog(null,
						"Empty user name not allowed.", "Error",
						JOptionPane.ERROR_MESSAGE);
				return;
			}
			InitializeClientMessage icm = new InitializeClientMessage(userName);
			Encoder enc = new Encoder();
			String mess;
			try {
				mess = enc.encode(icm);
				net.send(mess);
			} catch (JSONException e) {
				JOptionPane.showMessageDialog(null, "Can't connect to server!",
						"Warning", JOptionPane.ERROR_MESSAGE);
			}
			this.stage = 1;
			Tick();
		} catch (IOException e) {
			e.printStackTrace();
		}

	}

	public void Tick() {
		while (true) {
			if (this.stage == 1) {
				this.LH = new LobbyHandler(net, userName);
				StageFlipper passOn = LH.init(new StageFlipper());
				stop(passOn);
			} else if (this.stage == 20) {
				this.TH = new TiarHandler(net, userName);
				StageFlipper passOn = this.TH.init(lastMsg);
				stop(passOn);
			} else if (this.stage == 30) {
				this.SH = new SnakeHandler(net, userName);
				StageFlipper passOn = this.SH.init(lastMsg);
				stop(passOn);
			} else {
				System.out.println("Weird stage");
			}
		}
	}

	public void stop(StageFlipper flipper) {
		if (flipper.packageID == 102) {
			if (flipper.jm.GameType.equals("TicTacToe")) {
				this.lastMsg = flipper;
				this.stage = 20;
			} else if (flipper.jm.GameType.equals("Achtung Die Kurve")) {

				this.lastMsg = flipper;
				this.stage = 30;
			}

		} else if (flipper.packageID == 404) {
			this.stage = 1;
		} else {
			System.out
					.println("Unknown packageID for Monitor.stop(StageFlipper)");
		}
	}


}