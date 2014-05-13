package clientCore;

import java.io.IOException;

import javax.swing.JOptionPane;

import org.json.JSONException;

import packageManaging.InitializeClientMessage;
import packageManaging.InitializeClientMessageEncoder;



import clientHandlers.LobbyHandler;
import clientHandlers.TiarHandler;

import clientNetworking.NetManager;
import clientNetworking.Connection;


public class Monitor {
	
	public  Connection conn;
	public  NetManager net;
	public String userName;
	
	public Monitor() {
		
	}

	
	public void init(){
		
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
			
			setState(1);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
        

      
		
		
	}
	
	public void setState (int state){
		while(true){
		if(state == 1){
			LobbyHandler lh = new LobbyHandler(net, userName);
			state = lh.init();
			continue;
		}
		if(state == 2){
			 TiarHandler th = new TiarHandler(net, userName);
			 state = th.init();
		}
		else{
			System.out.println("hejdå");
		}
		}
		
	}
	
	
}