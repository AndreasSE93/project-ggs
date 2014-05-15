package clientCore;

import java.io.IOException;

import javax.swing.JOptionPane;

import org.json.JSONException;

import packageManaging.InitializeClientMessage;
import packageManaging.InitializeClientMessageEncoder;
import packageManaging.StageFlipper;



import clientHandlers.LobbyHandler;
import clientHandlers.TiarHandler;

import clientNetworking.NetManager;
import clientNetworking.Connection;


public class Monitor {
	
	public LobbyHandler LH;
	public TiarHandler TH;
	public  Connection conn;
	public  NetManager net;
	public String userName;
	public int stage;
	public StageFlipper lastMsg = null;
	
	public Monitor() {
		
	}

	
	public void init() {
		
		this.conn = new Connection("localhost", 8080);
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
				this.LH.init();
			} else if (this.stage == 20) {
				this.TH = new TiarHandler(net, userName);
				this.TH.init();
			} else if (this.stage == 10) {
				this.LH.loop = true;
				this.LH.runLobby();
			}
		}
	}
	
	public void stop(StageFlipper flipper) {
		
		if(flipper.packageID == 102) {
			this.LH.loop = false;
			this.lastMsg = flipper;
			this.stage = 20;
			
		} else if (flipper.packageID == 404) {
			this.stage = 10;
		} else {
			System.out.println("Unknown packageID for Monitor.stop(StageFlipper)");
		}
	}
	
	
	
}