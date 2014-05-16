package clientCore;

import java.io.IOException;

import javax.swing.JOptionPane;

import org.json.JSONException;

import packageManaging.InitializeClientMessage;
import packageManaging.InitializeClientMessageEncoder;
import packageManaging.StageFlipper;



import clientHandlers.LobbyHandler;
import clientHandlers.SnakeHandler;
import clientHandlers.TiarHandler;

import clientNetworking.NetManager;
import clientNetworking.Connection;


public class Monitor {
	
	public LobbyHandler LH;
	public TiarHandler TH;
	public SnakeHandler SH;
	public  Connection conn;
	public  NetManager net;
	public String userName;
	public int stage;
	public StageFlipper lastMsg = null;
	
	public Monitor() {
		
	}

	
	public void init() {
		
		this.conn = new Connection("130.243.137.236", 8080);
		this.net = new NetManager(conn);
        try {
			net.connectToServer();
			userName = (String)JOptionPane.showInputDialog("Write username!");
			if (userName == null) {
				return;
			} else if (userName.length() == 0) {
				JOptionPane.showMessageDialog(null, "Empty user name not allowed.", "Error", JOptionPane.ERROR_MESSAGE);
				return;
			}
			InitializeClientMessage icm = new InitializeClientMessage(userName);
			InitializeClientMessageEncoder icme = new InitializeClientMessageEncoder();
			String mess;
			try {
				mess = icme.encode(icm);
				net.send(mess);
			} catch (JSONException e) {
				JOptionPane.showMessageDialog(null, "Can't connect to server!", "Warning", JOptionPane.ERROR_MESSAGE);
			}
			this.LH = new LobbyHandler(net, userName);
			this.stage = 1;
			Tick();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
        		
		
	}
	
	public void Tick() {
		while(true) {
			if (this.stage == 1) {
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
		if(flipper.packageID == 102) {
			System.out.println(flipper.jm.GameType);
			if(flipper.jm.GameType.equals("TicTacToe")) {
				this.lastMsg = flipper;
				this.stage = 20;
			} else if (flipper.jm.GameType.equals("Achtung Die Kurve")) {
				System.out.println("STOP");
				this.lastMsg = flipper;
				this.stage = 30;
			}
			
		} else if (flipper.packageID == 404) {
			this.stage = 1;
		} else {
			System.out.println("Unknown packageID for Monitor.stop(StageFlipper)");
		}
	}
	
	
	
}